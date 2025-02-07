// 65070503403 Kritsanaphong Thaworana
public class AccountsReceivable {
    private final CustomerTransaction transaction;

    public AccountsReceivable(CustomerTransaction transaction) {
        if (transaction == null) {
            throw new IllegalArgumentException("Transaction cannot be null.");
        }
        this.transaction = transaction;
    }

    public void sendInvoice() {
        System.out.println("Invoice sent to " + transaction.getName());
    }

    public void postPayment() {
        System.out.println("Payment processed for " + transaction.getName());
    }

    // Override toString() for debugging
    @Override
    public String toString() {
        return "AccountsReceivable{transaction=" + transaction.getName() + "}";
    }
}
