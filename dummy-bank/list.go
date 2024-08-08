package dummybank

type Tool func()

var List = [6]Tool{
  CreateAccount,
  DepositMoney,
  WithdrawMoney,
  CheckBalance,
  ViewTransactionsHistory,
  Settings,
}
