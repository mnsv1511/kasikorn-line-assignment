package service

import (
	"fmt"
	"math"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/mnsv1511/kasikorn-line-assignment/constant"
	"github.com/mnsv1511/kasikorn-line-assignment/internal/core/service/domain"
)

func (s *ServiceImpl) GetDebitCardExample(c echo.Context, userId string) (*domain.GetDebitCardExampleResponse, error) {
	intUserId, err := strconv.Atoi(userId)
	if err != nil {
		log.Errorf("error req get debit card example: %s" ,err)
		return &domain.GetDebitCardExampleResponse{
			Status: &domain.StatusCode{
				Code:        string(constant.INVALID_REQUEST_CODE),
				Description: string(constant.INVALID_REQUEST_DES),
			},
		}, err
	}

	debitCardListData, err := s.repository.GetDebitCardsListByUserId(intUserId)
	if err != nil {
		log.Errorf("error repo get debit card list by user id: %s" ,err)
		return &domain.GetDebitCardExampleResponse{
			Status: &domain.StatusCode{
				Code:        string(constant.DEBIT_CARD_NOT_FOUND_CODE),
				Description: string(constant.DEBIT_CARD_NOT_FOUND_DES),
			},
		}, err
	}
	debitCardDesignListData, err := s.repository.GetDebitCardDesignListByUserId(intUserId)
	if err != nil {
		log.Errorf("error repo get debit card design list by user id: %s" ,err)
		return &domain.GetDebitCardExampleResponse{
			Status: &domain.StatusCode{
				Code:        string(constant.DEBIT_CARD_NOT_FOUND_CODE),
				Description: string(constant.DEBIT_CARD_NOT_FOUND_DES),
			},
		}, err
	}
	debitCardDetailListData, err := s.repository.GetDebitCardDetailListByUserId(intUserId)
	if err != nil {
		log.Errorf("error repo get debit card detail list by user id: %s" ,err)
		return &domain.GetDebitCardExampleResponse{
			Status: &domain.StatusCode{
				Code:        string(constant.DEBIT_CARD_NOT_FOUND_CODE),
				Description: string(constant.DEBIT_CARD_NOT_FOUND_DES),
			},
		}, err
	}
	debitCardStatusListData, err := s.repository.GetDebitCardStatusListByUserId(intUserId)
	if err != nil {
		log.Errorf("error repo get debit card status list by user id: %s" ,err)
		return &domain.GetDebitCardExampleResponse{
			Status: &domain.StatusCode{
				Code:        string(constant.DEBIT_CARD_NOT_FOUND_CODE),
				Description: string(constant.DEBIT_CARD_NOT_FOUND_DES),
			},
		}, err
	}
	debitCardListMap := make(map[int]domain.DebitCard)
	for _, debitCard := range debitCardListData {
		debitCardListMap[debitCard.ID] = domain.DebitCard{
			DebitCardId:   fmt.Sprintf("%d", debitCard.ID),
			DebitCardName: debitCard.Name,
		}
	}

	for _, debitCardDesign := range debitCardDesignListData {
		if debitCard, ok := debitCardListMap[debitCardDesign.CardID]; ok {
			debitCard.DebitCardColor = debitCardDesign.Color
			debitCard.DebitCardBorderColor = debitCardDesign.BorderColor
			debitCardListMap[debitCardDesign.CardID] = debitCard
		}
	}

	for _, debitCardDetail := range debitCardDetailListData {
		if debitCard, ok := debitCardListMap[debitCardDetail.CardID]; ok {
			debitCard.DebitCardIssuer = debitCardDetail.Issuer
			debitCard.DebitCardNumber = debitCardDetail.Number
			debitCardListMap[debitCardDetail.CardID] = debitCard
		}
	}

	for _, debitCardStatus := range debitCardStatusListData {
		if debitCard, ok := debitCardListMap[debitCardStatus.CardID]; ok {
			debitCard.DebitCardStatus = debitCardStatus.Status
			debitCardListMap[debitCardStatus.CardID] = debitCard
		}
	}

	var debitCardResp []domain.DebitCard
	for _, debitCard := range debitCardListData {
		debitCardResp = append(debitCardResp, debitCardListMap[debitCard.ID])
	}
	lenOfExampleDebitCard := 4.0
	minOfDebitCard := int(math.Floor(math.Min(lenOfExampleDebitCard, float64(len(debitCardResp)))))
	return &domain.GetDebitCardExampleResponse{
		DebitCard: debitCardResp[:minOfDebitCard],
	}, nil
}
