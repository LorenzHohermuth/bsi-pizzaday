# BSI Pizzaday List Generator
## Install
- dependencies: [`go`](https://go.dev/)
- `$ go install github.com/LorenzHohermuth/bsi-pizzaday@latest`
- create `.config` direcotry
- set enviroment variable `DOT_DIR` to `.config` directory
- create file in `.config` directory named `.pizzaconfig`
- configure it this is the default for Zurich
```
pizza.vegi = [Vegi, Funghi]
pizza.meat = [Antonio, Tonno, Emiliana, Prosciutto, Calabrese]

bsi.location = Zürich

available.slots = [Slot 1: 11:45, Slot 2: 12:15]
available.pizzaTypes = [Pizza mit Fleisch, Vegetarische Pizza]

slot1.timeToPickUp = Slot 1: 11:30
slot2.timeToPickUp = Slot 2: 12:00

```

## Run App
- Download Excel file from Story
- Open it
- File > Save As > Change file type to CSV  + Enter name > Save
- `$ bsi-pizzaday generate --file .\path\to\csv\file`
- there you go
