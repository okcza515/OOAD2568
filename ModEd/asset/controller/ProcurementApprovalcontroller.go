// MEP-1014
package controller

import (
	model "ModEd/asset/model"

	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProcurementApprovalController struct {
	db *gorm.DB
}

func (ctrl *ProcurementApprovalController) ListApprovals(c *gin.Context) {
	var approvals []model.ProcurementApproval

	if err := ctrl.db.Unscoped().Find(&approvals).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch approvals"})
		return
	}

	var result []gin.H
	for _, a := range approvals {
		entry := gin.H{
			"procurement_approval_id": a.ProcurementApprovalID,
			"description":             a.Description,
			"status":                  a.Status,
			"approver_id":             a.ApproverID,
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

	if err := ctrl.db.First(&approval, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Approval not found"})
		return
	}

	c.JSON(http.StatusOK, approval)
}

func (ctrl *ProcurementApprovalController) FilterApprovals(c *gin.Context) {
	var approvals []model.ProcurementApproval
	approverID := c.Query("approver_id")
	status := c.Query("status")

	query := ctrl.db.Model(&model.ProcurementApproval{})

	if approverID != "" {
		query = query.Where("approver_id = ?", approverID)
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

	input.Status = strings.ToLower(input.Status)

	var approval model.ProcurementApproval
	if err := ctrl.db.First(&approval, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Approval not found"})
		return
	}

	approval.Status = input.Status
	if input.Status == "Approved" {
		now := time.Now()
		approval.ApprovalTime = &now
	}

	if err := ctrl.db.Save(&approval).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update status"})
		return
	}

	c.JSON(http.StatusOK, approval)
}

func (ctrl *ProcurementApprovalController) DeleteApproval(c *gin.Context) {
	id := c.Param("id")

	if err := ctrl.db.Delete(&model.ProcurementApproval{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete approval"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Approval deleted"})
}

func (ctrl *ProcurementApprovalController) OnApproved(id uint) error {
	var approval model.ProcurementApproval
	if err := ctrl.db.First(&approval, id).Error; err != nil {
		return err
	}
	approval.Status = "approved"
	now := time.Now()
	approval.ApprovalTime = &now
	return ctrl.db.Save(&approval).Error
}

func (ctrl *ProcurementApprovalController) OnRejected(id uint) error {
	return ctrl.db.Model(&model.ProcurementApproval{}).
		Where("procurement_approval_id = ?", id).
		Update("status", "rejected").Error
}
