// MEP-1014
package controller

import (
	model "ModEd/asset/model/Procurement"

	"gorm.io/gorm"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
)

type ProcurementApprovalController struct {
	DB *gorm.DB
}

func (ctrl *ProcurementApprovalController) ListApprovals(c *gin.Context) {
	var approvals []model.ProcurementApproval

	if err := ctrl.DB.Unscoped().Find(&approvals).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch approvals"})
		return
	}

	var result []gin.H
	for _, a := range approvals {
		entry := gin.H{
			"procurement_approval_id": a.ProcurementApprovalID,
			"description":             a.Description,
			"status":                  a.Status,
			"approvers_id":           a.ApproversID,
		}

		if a.DeletedAt.Valid {
			entry["deleted_at"] = a.DeletedAt.Time
		} else if a.ApprovalTime != nil {
			entry["approval_time"] = a.ApprovalTime
		}

		result = append(result, entry)
	}

	c.JSON(http.StatusOK, result)
}

func (ctrl *ProcurementApprovalController) GetApprovalByID(c *gin.Context) {
	id := c.Param("id")
	var approval model.ProcurementApproval

	if err := ctrl.DB.First(&approval, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Approval not found"})
		return
	}

	c.JSON(http.StatusOK, approval)
}

func (ctrl *ProcurementApprovalController) FilterApprovals(c *gin.Context) {
	var approvals []model.ProcurementApproval
	approverID := c.Query("approver_id")
	status := c.Query("status")

	query := ctrl.DB.Model(&model.ProcurementApproval{})

	if approverID != "" {
		query = query.Where("approvers_id = ?", approverID)
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}

	if err := query.Find(&approvals).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to filter approvals"})
		return
	}

	c.JSON(http.StatusOK, approvals)
}

func (ctrl *ProcurementApprovalController) UpdateApprovalStatus(c *gin.Context) {
	id := c.Param("id")
	var input struct {
		Status string `json:"status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var approval model.ProcurementApproval
	if err := ctrl.DB.First(&approval, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Approval not found"})
		return
	}

	approval.Status = input.Status
	if input.Status == "Approved" {
		now := time.Now()
		approval.ApprovalTime = &now
	}

	if err := ctrl.DB.Save(&approval).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update status"})
		return
	}

	c.JSON(http.StatusOK, approval)
}

func (ctrl *ProcurementApprovalController) DeleteApproval(c *gin.Context) {
	id := c.Param("id")

	if err := ctrl.DB.Delete(&model.ProcurementApproval{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete approval"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Approval deleted"})
}
