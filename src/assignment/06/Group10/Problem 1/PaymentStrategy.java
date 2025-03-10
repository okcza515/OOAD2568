//65070501078 Nawaphon Promnan
public interface PaymentStrategy {
    public void collectPaymentDetails();
    public boolean pay(int amount);
}