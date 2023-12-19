package usecase

import (
	"fmt"
	addressRep "server/internal/Address/repository"
	sessionRep "server/internal/Session/repository"
	"server/internal/domain/dto"
	"server/internal/domain/entity"
)

//AddressUsecaseI interface
type AddressUsecaseI interface {
	CreateAddress(reqAddress *dto.ReqCreateAddress) error
	DeleteAddress(id uint, sessionToken string) error
	//GetAddresses(UserID uint) (*dto.RespGetAddresses, error)
	SetAddress(id uint, sessionToken string) error
}

//AddressUsecase struct
type AddressUsecase struct {
	addressRepo addressRep.AddressRepositoryI
	sessionRepo sessionRep.SessionRepositoryI
}

//NewAddressUsecase creates address usecase
func NewAddressUsecase(addressRepI addressRep.AddressRepositoryI, sessionRepI sessionRep.SessionRepositoryI) *AddressUsecase {
	return &AddressUsecase{
		addressRepo: addressRepI,
		sessionRepo: sessionRepI,
	}
}

//CreateAddress creates address
func (ad *AddressUsecase) CreateAddress(reqAddress *dto.ReqCreateAddress) error {
	address := dto.ToEntityCreateAddress(reqAddress)

	if len(address.City) == 0 || len(address.Street) == 0 || len(address.House) == 0 {
		return entity.ErrBadRequest
	}

	cookie, err := ad.sessionRepo.Check(reqAddress.Cookie)
	if err != nil {
		return err
	}
	if cookie == nil {
		return entity.ErrNotFound
	}

	addresses, err := ad.addressRepo.GetAddresses(cookie.UserID)
	if err != nil {
		fmt.Println("get", err)
		return err
	}
	fmt.Println("")
	if addresses != nil {
		for _, checkAd := range addresses.Addresses {
			if checkAd.City == address.City && checkAd.Street == address.Street && checkAd.House == address.House && checkAd.Flat == address.Flat {
				return entity.ErrAddressAlreadyExist
			}
		}
	}
	return ad.addressRepo.CreateAddress(dto.ToDBCreateAddress(address, cookie.UserID))
}

//DeleteAddress deletes address
func (ad *AddressUsecase) DeleteAddress(id uint, sessionToken string) error {
	cookie, err := ad.sessionRepo.Check(sessionToken)
	if err != nil {
		return err
	}
	if cookie == nil {
		return entity.ErrNotFound
	}

	address := &dto.DBReqDeleteUserAddress{
		UserID:    cookie.UserID,
		AddressID: id,
	}

	return ad.addressRepo.DeleteAddress(address)
}

//SetAddress sets address
func (ad *AddressUsecase) SetAddress(id uint, sessionToken string) error {

	cookie, err := ad.sessionRepo.Check(sessionToken)
	if err != nil {
		return err
	}
	if cookie == nil {
		return entity.ErrNotFound
	}

	address := &dto.DBReqUpdateUserAddress{
		UserID:    cookie.UserID,
		AddressID: id,
	}

	return ad.addressRepo.SetAddress(address)
}
