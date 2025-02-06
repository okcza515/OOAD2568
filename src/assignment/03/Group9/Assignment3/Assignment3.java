// 6. Service layer to orchestrate the operations
//Wat
class TransactionService {
    private final TransactionGenerator transactionGenerator;
    private final AccountsReceivable accountsReceivable;

    public TransactionService(TransactionGenerator transactionGenerator, 
                            AccountsReceivable accountsReceivable) {
        this.transactionGenerator = transactionGenerator;
        this.accountsReceivable = accountsReceivable;
    }

    public void processTransaction(Customer customer, Product product) 
            throws TransactionException, AccountingException {
        Transaction transaction = transactionGenerator.generateTransaction(customer, product);
        transaction.execute();
        accountsReceivable.recordTransaction(transaction);
    }

    public BigDecimal getCustomerBalance(Customer customer) {
        return accountsReceivable.getCustomerBalance(customer);
    }
}

// 7. Custom exceptions
class TransactionException extends Exception {
    public TransactionException(String message) {
        super(message);
    }
}

class AccountingException extends Exception {
    public AccountingException(String message) {
        super(message);
    }
}