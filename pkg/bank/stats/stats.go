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
	 