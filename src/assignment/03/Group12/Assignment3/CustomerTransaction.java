import java.util.Date;
import java.util.List;

public class CustomerTransaction {

	private List<Product> products;
	private Customer customer;

	public CustomerTransaction(Customer customer, List<Product> products) {
		this.products = products;
		this.customer = customer;
	}

	public String getName() {
		return "name";
	}

	public Date getDate() {
		return new Date();
	}

	public String productBreakDown() {
		return "list of products for reporting";
	}

	public void prepareInvoice() {
		System.out.println("invoice prepared...");
	}

	public void chargeCustomer() {
		System.out.println("charged the customer");
	}
}
