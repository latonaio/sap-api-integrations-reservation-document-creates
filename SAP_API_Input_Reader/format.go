package sap_api_input_reader

import (
	"sap-api-integrations-reservation-document-creates/SAP_API_Caller/requests"
)

func (sdc *SDC) ConvertToHeader() *requests.Header {
	data := sdc.Reservation
	return &requests.Header{
		Reservation:                  data.Reservation,
		OrderID:                      data.OrderID,
		GoodsMovementType:            data.GoodsMovementType,
		CostCenter:                   data.CostCenter,
		GoodsRecipientName:           data.GoodsRecipientName,
		ReservationDate:              data.ReservationDate,
		Customer:                     data.Customer,
		WBSElement:                   data.WBSElement,
		ControllingArea:              data.ControllingArea,
		SalesOrder:                   data.SalesOrder,
		SalesOrderItem:               data.SalesOrderItem,
		SalesOrderScheduleLine:       data.SalesOrderScheduleLine,
		AssetNumber:                  data.AssetNumber,
		AssetSubNumber:               data.AssetSubNumber,
		NetworkNumberForAcctAssgmt:   data.NetworkNumberForAcctAssgmt,
		IssuingOrReceivingPlant:      data.IssuingOrReceivingPlant,
		IssuingOrReceivingStorageLoc: data.IssuingOrReceivingStorageLoc,
	}
}

func (sdc *SDC) ConvertToItem() *requests.Item {
	dataReservation := sdc.Reservation
	data := sdc.Reservation.ReservationItem
	return &requests.Item{
		Reservation:                    dataReservation.Reservation,
		ReservationItem:                data.ReservationItem,
		RecordType:                     data.RecordType,
		Product:                        data.Product,
		RequirementType:                data.RequirementType,
		MatlCompRequirementDate:        data.MatlCompRequirementDate,
		Plant:                          data.Plant,
		ManufacturingOrderOperation:    data.ManufacturingOrderOperation,
		GoodsMovementIsAllowed:         data.GoodsMovementIsAllowed,
		StorageLocation:                data.StorageLocation,
		Batch:                          data.Batch,
		DebitCreditCode:                data.DebitCreditCode,
		BaseUnit:                       data.BaseUnit,
		GLAccount:                      data.GLAccount,
		GoodsMovementType:              data.GoodsMovementType,
		EntryUnit:                      data.EntryUnit,
		QuantityIsFixed:                data.QuantityIsFixed,
		CompanyCodeCurrency:            data.CompanyCodeCurrency,
		IssuingOrReceivingPlant:        data.IssuingOrReceivingPlant,
		IssuingOrReceivingStorageLoc:   data.IssuingOrReceivingStorageLoc,
		PurchasingDocument:             data.PurchasingDocument,
		PurchasingDocumentItem:         data.PurchasingDocumentItem,
		Supplier:                       data.Supplier,
		ResvnItmRequiredQtyInBaseUnit:  data.ResvnItmRequiredQtyInBaseUnit,
		ReservationItemIsFinallyIssued: data.ReservationItemIsFinallyIssued,
		ReservationItmIsMarkedForDeltn: data.ReservationItmIsMarkedForDeltn,
		ResvnItmRequiredQtyInEntryUnit: data.ResvnItmRequiredQtyInEntryUnit,
		ResvnItmWithdrawnQtyInBaseUnit: data.ResvnItmWithdrawnQtyInBaseUnit,
		ResvnItmWithdrawnAmtInCCCrcy:   data.ResvnItmWithdrawnAmtInCCCrcy,
		GoodsRecipientName:             data.GoodsRecipientName,
		UnloadingPointName:             data.UnloadingPointName,
		ReservationItemText:            data.ReservationItemText,
	}
}
