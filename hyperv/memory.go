package hyperv

import (
	"net/http"
	"wmi-rest/utilities"

	"github.com/gin-gonic/gin"
)

func Memory(c *gin.Context) {
	input := c.Param("machid")

	if input == "" {
		c.Data(returnResponse("No VM ID specified", http.StatusBadRequest, "failure", "error"))
		return
	}

	if input == "all" {
		output, err := utilities.CommandLine(`Get-WmiObject -Namespace 'root\virtualization\v2' -Class Msvm_MemorySettingData -Filter "Caption like 'Memory'" | Select-Object -Property InstanceID, VirtualQuantity | ConvertTo-Json`)
		if err != nil {
			c.Data(returnResponse(err.Error(), http.StatusInternalServerError, "failure", "error"))
			return
		}
		c.Data(returnResponse(output, http.StatusOK, "success", "Memory info is displayed in data field"))
		return
	}

	output, err := utilities.CommandLine(`Get-VM -Id ` + input + ` | Get-VMMemory | Select-Object -Property InstanceID, VirtualQuantity | ConvertTo-Json`)
	if err != nil {
		c.Data(returnResponse(err.Error(), http.StatusInternalServerError, "failure", "error"))
		return
	}
	c.Data(returnResponse(output, http.StatusOK, "success", "Memory info is displayed in data field"))
}
