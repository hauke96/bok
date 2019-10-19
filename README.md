# bok
This simple accounting tool called bok (short for icelandic "bókhald") keeps track of bills, receips and expenses.

## The idea
You enter date, amount, description and category for each bill, receip, expense, etc. which are saved in a **store** file (usually `account.json`).

To use the entries in LibreOffice Calc (or Excel but no guarantee that it works) you can also export the store into a CSV file.

# Building
1. Clone or download this repo
2. Execute `make`

Done :D

# Usage

## Workflow
Let's run `./bok` and add an entry by typing `a` in the REPL popping up.

```
foo@bar $ ./bok
Welcome to bok
> a
Add new entry:
  Date: 2019-10-18
  Amount: 12,99
  Description: Book
  Category: 3
```

To save just type `w` (→ write).

```
> w
Saved
```

To quit bok type `q`.

```
> q
Bye
```

## Export
To use the data in e.g. LibreOffice Calc, you can export everything into a CSV file:
```
foo@bar $ ./bok export
Exporting finished succesfully
  Format: csv
  File: account_export.csv
```
The resulting CSV file (with this one entry) look like this:
```
"12,99","Book","18.10.2019","3"
```
