
public class Order {
	private int totalCost = 0;
	private boolean isClosed = false;

	public void processOrder(PaymentStrategy strategy) {
		strategy.collectPaymentDetails();
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
// 65070501048 Rattipong Sakunjeen