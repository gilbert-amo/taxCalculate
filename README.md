
# taxCalculate 🧮

A command-line utility for calculating taxes on prices, supporting both inclusive and exclusive tax calculations. Built with Go.

## Features ✨

- ✅ Calculate multiple taxes simultaneously
- 🔄 Handle both tax-inclusive and tax-exclusive prices
- 📊 Detailed tax breakdown report
- 🛠️ Simple command-line interface
- 💰 GHS currency formatting


## Installation ⚙️


```bash
go get github.com/gilbert-amo/taxCalculate
```
    
## Documentation

### Usage
#### Basic Import

```bash

import (
    "github.com/gilbert-amo/taxCalculate/tax"
)

```

Calculating for a tax 
 ```bash 

var taxes []tax.Tax

subtotal, total, taxAmounts := tax.CalculateTotal(price, taxes, isInclusive)

 ```

