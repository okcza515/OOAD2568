public class Order {
	private int totalCost = 0;
	private boolean isClosed = false;

	public void processOrder(PaymentStrategy strategy) {
			strategy.collectPaymentDetails();
	}

	public boolean pay(PaymentStrategy strategy) {
			return strategy.pay(totalCost);
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
