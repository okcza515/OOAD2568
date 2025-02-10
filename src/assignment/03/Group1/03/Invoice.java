public class Invoice {
    private CustomerTransaction transaction;

    public Invoice(CustomerTransaction transaction) {
        this.transaction = transaction;
    }

    public CustomerTransaction getTransaction() {
        return transaction;
    }
    
    @Override
    public String toString() {
        return "Invoice for " + transaction.getCustomer().getName() +
               " on " + transaction.getDate() +
               " with " + transaction.getProducts().size() + " product(s)";
    }
}
