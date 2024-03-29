package stats

import (
	"github.com/hirs-500/bank/pkg/bank/types"
)

// Avg рассчитывает среднюю сумму платежа 
func Avg (payments []types.Payment)  types.Money {
	 var total types.Money
	 payment := len(payments)
	 for _, v := range payments {
		 total+= v.Amount}
		 
		 return total/types.Money(payment)
		 
	 }
	 //  TotalInCategory - находить сумму пакупок определённой котегории
	 func TotalInCategory(payments []types.Payment, category types.Category ) types.Money {
		var categorySum types.Money
		for _, v := range payments {
			if v.Category == category{
			   categorySum+=v.Amount
			}}
			
			return categorySum
		
		
	}