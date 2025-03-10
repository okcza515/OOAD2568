//65070501008
public interface PaymentStrategy {
    void collectPaymentDetails();
    boolean pay(int paymentAmount);
}
