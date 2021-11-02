package stats

import (
	"fmt"

	"github.com/DariaYudina004/bank/pkg/bank/types"
)

func ExampleAvg()  {
	payments := []types.Payment{
		{
			Amount: 1000_00,
		},
		{
			Amount: 2000_00,
		},
		{
			Amount: 2000_00,
		},
		{
			Amount: 3000_00,
		},

		
	}
	fmt.Println(Avg(payments))
	//Output:
	//200000

}