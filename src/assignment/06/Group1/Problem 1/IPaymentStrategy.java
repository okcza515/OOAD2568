public interface IPaymentStrategy {
    void collectPaymentDetails();
  
    boolean pay(int amount);
}