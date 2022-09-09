package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-reservation-document-creates/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) (*Header, error) {
	pm := &responses.Header{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	header := &Header{
		Reservation:                   data.Reservation,
		OrderID:                       data.OrderID,
		GoodsMovementType:             data.GoodsMovementType,
		CostCenter:                    data.CostCenter,
		GoodsRecipientName:            data.GoodsRecipientName,
		ReservationDate:               data.ReservationDate,
		Customer:                      data.Customer,
		WBSElement:                    data.WBSElement,
		ControllingArea:               data.ControllingArea,
		SalesOrder:                    data.SalesOrder,
		SalesOrderItem:                data.SalesOrderItem,
		SalesOrderScheduleLine:        data.SalesOrderScheduleLine,
		AssetNumber:                   data.AssetNumber,
		AssetSubNumber:                data.AssetSubNumber,
		NetworkNumberForAcctAssgmt:    data.NetworkNumberForAcctAssgmt,
		IssuingOrReceivingPlant:       data.IssuingOrReceivingPlant,
		IssuingOrReceivingStorageLoc:  data.IssuingOrReceivingStorageLoc,
	}
	
	return header, nil
}

func ConvertToItem(raw []byte, l *logger.Logger) (*Item, error) {
	pm := &responses.Item{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Item. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	item := &Item{
        Reservation:                   data.Reservation,
		ReservationItem:               data.ReservationItem,
		RecordType:                    data.RecordType,
		Product:                       data.Product,
		RequirementType:               data.RequirementType,
		MatlCompRequirementDate:       data.MatlCompRequirementDate,
		Plant:                         data.Plant,
		ManufacturingOrderOperation:   data.ManufacturingOrderOperation,
		GoodsMovementIsAllowed:        data.GoodsMovementIsAllowed,
		StorageLocation:               data.StorageLocation,
		Batch:                         data.Batch,
		DebitCreditCode:               data.DebitCreditCode,
		BaseUnit:                      data.BaseUnit,
		GLAccount:                     data.GLAccount,
		GoodsMovementType:             data.GoodsMovementType,
		EntryUnit:                     data.EntryUnit,
		QuantityIsFixed:               data.QuantityIsFixed,
		CompanyCodeCurrency:           data.CompanyCodeCurrency,
		IssuingOrReceivingPlant:       data.IssuingOrReceivingPlant,
		IssuingOrReceivingStorageLoc:  data.IssuingOrReceivingStorageLoc,
		PurchasingDocument:            data.PurchasingDocument,
		PurchasingDocumentItem:        data.PurchasingDocumentItem,
		Supplier:                      data.Supplier,
		ResvnItmRequiredQtyInBaseUnit: data.ResvnItmRequiredQtyInBaseUnit,
		ReservationItemIsFinallyIssued: data.ReservationItemIsFinallyIssued,
		ReservationItmIsMarkedForDeltn: data.ReservationItmIsMarkedForDeltn,
		ResvnItmRequiredQtyInEntryUnit: data.ResvnItmRequiredQtyInEntryUnit,
		ResvnItmWithdrawnQtyInBaseUnit: data.ResvnItmWithdrawnQtyInBaseUnit,
		ResvnItmWithdrawnAmtInCCCrcy:  data.ResvnItmWithdrawnAmtInCCCrcy,
		GoodsRecipientName:            data.GoodsRecipientName,
		UnloadingPointName:            data.UnloadingPointName,
		ReservationItemText:           data.ReservationItemText,
	}

	return item, nil
}
