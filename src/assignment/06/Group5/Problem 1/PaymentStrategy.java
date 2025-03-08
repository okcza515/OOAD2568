public interface PaymentStrategy {
    void collectPaymentDetails();
    boolean pay(int amount);
}
// 65070501048 Rattipong Sakunjeen