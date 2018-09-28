# Exercise #20: Building PDFs

## Exercise details

There are two goals with this exercise.

1. Create a PDF invoice given some data about a fake customer's recent transactions
2. Generate a course completion certificate and add your name to it!

**NOTE:** *The design part of creating PDFs can take a while, so don't get upset if this exercise takes you a few hours to get all those details right.*

### Creating a PDF invoice

Given data like the following, generate a PDF invoice.

```go
data := []struct {
  UnitName       string
  PricePerUnit   int
  UnitsPurchased int
}{
  {
    UnitName:       "2x6 Lumber - 8'",
    PricePerUnit:   375, // in cents
    UnitsPurchased: 220,
  }, {
    UnitName:       "Drywall Sheet",
    PricePerUnit:   822, // in cents
    UnitsPurchased: 50,
  }, {
    UnitName:       "Paint",
    PricePerUnit:   1455, // in cents
    UnitsPurchased: 3,
  },
}
```

You may not do this all dynamically at first, but once you have your code working try to clean it up and make it easier to create dynamic invoices as the data you are provided changes. You can even try to come up with a way to add support for multiple pages.


### Course Completion Certificate

Create a course completion certificate that dynamically inserts your name.

### Notes

**I will be using the `gofpdf` package**

You can find it here: <https://github.com/jung-kurt/gofpdf>

The docs are on godoc: <https://godoc.org/github.com/jung-kurt/gofpdf>

The library can take some time to get used to, especially some of the string arguments like alignment, fill, etc. It can also be confusing learning which function to use for text etc as each has its own pros and cons. Just experiment and a bit and try to get *something* working 🙂


## Bonus

Make the invoice and the course completion certificate look better, and make the invoice work for multiple pages of data.
