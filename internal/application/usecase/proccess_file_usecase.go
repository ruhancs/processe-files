package usecase

import (
	"bytes"
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

func (usecase *ProccessFileUseCase) Execute(dataFile multipart.File) (string, error) {
	var file = new(bytes.Buffer)
	_, err := io.Copy(file, dataFile)
	if err != nil {
		return "error to procces the file", err
	}
	defer dataFile.Close()
	data := file.Bytes()
	lines := strings.Split(string(data), "\n")

	userMap := make(map[string]int)
	for _, line := range lines {
		controlProcces := len(lines)
		transactionType, err := strconv.Atoi(line[0:1])
		transactionDate := line[1:26]
		productName := line[26:56]
		seller := line[66:86]
		value, err := strconv.Atoi(line[56:66])
		if err != nil {
			return "invalid transaction value", err
		}
		user, err := usecase.UserRepository.FindByName(seller)
		if err != nil {
			return "seller not registered", err
		}
		_, err = usecase.ProductRepository.Get(productName)
		if err != nil {
			return "product not registered", err
		}
		usecase.checkMap(userMap, seller, user.Balance)
		if controlProcces == 0 {
			return "ok", nil
		}
		controlProcces--
		switch transactionType {
		case 1:
			userMap[seller] += value
			err := usecase.createTransaction(transactionType, value, transactionDate, productName, seller)
			if err != nil {
				return "error to save transaction", err
			}
			continue
		case 2:
			msg, err := usecase.proccesAffiliateSale(productName, value, userMap)
			if err != nil {
				return msg, err
			}
			err = usecase.createTransaction(transactionType, value, transactionDate, productName, seller)
			if err != nil {
				return "error to save transaction", err
			}
			continue
		case 3:
			msg, err := usecase.commissionPayment(productName, seller, value, userMap)
			if err != nil {
				return msg, err
			}
			err = usecase.createTransaction(transactionType, value, transactionDate, productName, seller)
			if err != nil {
				return "error to save transaction", err
			}
			continue
		}
	}

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

func (usecase *ProccessFileUseCase) proccesAffiliateSale(productName string, value int, usermap map[string]int) (string, error) {
	fmt.Println("proccess affiliate sale")
	fmt.Println(usermap)
	product, err := usecase.ProductRepository.Get(productName)
	if err != nil {
		return "product not registered", err
	}
	usermap[product.ProducerName] += value
	fmt.Println("MAP UPDATED")
	fmt.Println(usermap)
	return "", nil
}

func (usecase *ProccessFileUseCase) commissionPayment(productName, seller string, value int, usermap map[string]int) (string, error) {
	fmt.Println("proccess affiliate commission payment")
	fmt.Println(usermap)
	product, err := usecase.ProductRepository.Get(productName)
	if err != nil {
		return "product not registered", err
	}
	usermap[seller] += value
	usermap[product.ProducerName] -= value
	fmt.Println("MAP UPDATED")
	fmt.Println(usermap)
	return "", nil
}

func (usecase *ProccessFileUseCase) createTransaction(transactionType, value int, date, productName, seller string) error {
	transation, err := entity.NewTransaction(transactionType, value, date, productName, seller)
	if err != nil {
		return err
	}
	err = usecase.TransactionRepository.Create(transation)
	if err != nil {
		return err
	}
	return nil
}

func(usecase *ProccessFileUseCase) updateUserBalance(usersMap map[string]int) error {
	for user,value := range usersMap {
		err := usecase.UserRepository.UpdateBalance(user,value)
		if err != nil {
			return err
		}
	}
	return nil
}
