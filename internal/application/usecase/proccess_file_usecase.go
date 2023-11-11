package usecase

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"strconv"
	"strings"

	"github.com/ruhancs/hubla-test/internal/domain/entity"
	"github.com/ruhancs/hubla-test/internal/domain/gateway"
)

type ProccessFileUseCase struct {
	TransactionRepository gateway.TransactionRepositoryInterface
	ProductRepository     gateway.ProductRepositoryInterface
	UserRepository        gateway.UserRepositoryInterface
}

func NewProccessFileUseCase(
	transactionRepository gateway.TransactionRepositoryInterface,
	productRepository gateway.ProductRepositoryInterface,
	userRepository gateway.UserRepositoryInterface,
) *ProccessFileUseCase {
	return &ProccessFileUseCase{
		TransactionRepository: transactionRepository,
		ProductRepository:     productRepository,
		UserRepository:        userRepository,
	}
}

func (usecase *ProccessFileUseCase) Execute(ctx context.Context,dataFile multipart.File) (string, error) {
	var file = new(bytes.Buffer)
	_, err := io.Copy(file, dataFile)
	if err != nil {
		return "error to procces the file", err
	}
	defer dataFile.Close()
	data := file.Bytes()
	lines := strings.Split(string(data), "\n")

	userMap := make(map[string]int)
	countLines := 0
	for _, line := range lines {
		if len(line) != 0 {
			countLines ++
			transactionType, err := strconv.Atoi(line[0:1])
			transactionDate := line[1:26]
			productName := strings.TrimSpace(line[26:56])
			endLine := len(line)
			seller := line[66:endLine]
			value, err := strconv.Atoi(line[56:66])
			if err != nil {
				return "invalid transaction value", err
			}
			user, err := usecase.UserRepository.FindByName(ctx,seller)
			if err != nil {
				fmt.Println("error to get user")
				return "seller not registered", err
			}
			_, err = usecase.ProductRepository.Get(ctx,productName)
			if err != nil {
				fmt.Println("error to get product")
				return "product not registered", err
			}
			usecase.checkMap(userMap, seller, user.Balance)
			
			switch transactionType {
			case 1:
				userMap[seller] += value
				err := usecase.createTransaction(ctx,transactionType, value, transactionDate, productName, seller)
				if err != nil {
					return "error to save transaction", err
				}
				continue
			case 2:
				msg, err := usecase.proccesAffiliateSale(ctx,productName, value, userMap)
				if err != nil {
					return msg, err
				}
				fmt.Println(countLines)
				err = usecase.createTransaction(ctx,transactionType, value, transactionDate, productName, seller)
				if err != nil {
					return "error to save transaction", err
				}
				continue				
			case 3:
				err = usecase.createTransaction(ctx,transactionType, value, transactionDate, productName, seller)
				if err != nil {
					return "error to save transaction", err
				}
				continue				
			case 4:
				msg, err := usecase.commissionPayment(ctx,productName, seller, value, userMap)
				if err != nil {
					return msg, err
				}
				fmt.Println(countLines)
				err = usecase.createTransaction(ctx,transactionType, value, transactionDate, productName, seller)
				if err != nil {
					return "error to save transaction", err
				}
				continue
			}
		}
	}

	fmt.Println("map balance final")
	fmt.Println(userMap)
	err = usecase.updateUserBalance(ctx,userMap)
	if err != nil {
		return "error to update users balance",err
	}
	userMap = make(map[string]int)
	return "Process all transactions successfuly", nil
}

func (usecase *ProccessFileUseCase) checkMap(usersMap map[string]int, userName string, userBalance int) {
	_, ok := usersMap[userName]
	if !ok {
		usersMap[userName] = userBalance
	}
	return
}

func (usecase *ProccessFileUseCase) updateBalanceOnMap(usersMap map[string]int, seller string, value int) (string, error) {
	usersMap[seller] += value
	return "", nil
}

func (usecase *ProccessFileUseCase) proccesAffiliateSale(ctx context.Context,productName string, value int, usermap map[string]int) (string, error) {
	product, err := usecase.ProductRepository.Get(ctx ,productName)
	if err != nil {
		fmt.Println("error to find product in proccesAffiliateSale")
		return "product not registered", err
	}
	usermap[product.ProducerName] += value
	return "", nil
}

func (usecase *ProccessFileUseCase) commissionPayment(ctx context.Context,productName, seller string, value int, usermap map[string]int) (string, error) {
	product, err := usecase.ProductRepository.Get(ctx,productName)
	if err != nil {
		fmt.Println("error to find product in commissionPayment")
		return "product not registered", err
	}
	usermap[seller] += value
	usermap[product.ProducerName] -= value
	return "", nil
}

func (usecase *ProccessFileUseCase) createTransaction(ctx context.Context,transactionType, value int, date, productName, seller string) error {
	transation, err := entity.NewTransaction(transactionType, value, date, productName, seller)
	if err != nil {
		return err
	}
	err = usecase.TransactionRepository.Create(ctx,transation)
	if err != nil {
		return err
	}
	return nil
}

func(usecase *ProccessFileUseCase) updateUserBalance(ctx context.Context,usersMap map[string]int) error {
	//fmt.Println(usersMap)
	for user,value := range usersMap {
		err := usecase.UserRepository.UpdateBalance(ctx,user,value)
		if err != nil {
			return err
		}
	}
	return nil
}
