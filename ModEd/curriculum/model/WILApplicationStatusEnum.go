package model

type WILApplicationStatusEnum string

const (
    WIL_APP_PENDING  WILApplicationStatusEnum = "Pending"
    WIL_APP_APPROVED WILApplicationStatusEnum = "Approved"
    WIL_APP_REJECTED WILApplicationStatusEnum = "Rejected"
)