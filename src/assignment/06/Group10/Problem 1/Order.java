//65070501078 Nawaphon Promnan
public class Order {
    private int totalCost = 0;
    private boolean isClosed = false;
    private PaymentStrategy paymentStrategy;

    public void processOrder(PaymentStrategy strategy) {
        this.paymentStrategy = strategy; 
        strategy.collectPaymentDetails();
    }

    public boolean pay() {
        return paymentStrategy != null && paymentStrategy.pay(totalCost);
    }

    public void setTotalCost(int cost) {
        this.totalCost += cost;
    }

    public int getTotalCost() {
        return totalCost;
    }

    public boolean isClosed() {
        return isClosed;
    }

    public void setClosed() {
        isClosed = true;
    }
}