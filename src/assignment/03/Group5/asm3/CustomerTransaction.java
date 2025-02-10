import java.util.Date;
import java.util.List;

public class CustomerTransaction {

	private List<Product> products;
	private Customer customer;
	private Date transactionDate;

	public CustomerTransaction(Customer customer, List<Product> products) {
		this.products = products;
		this.customer = customer;
		this.transactionDate = new Date();
	}

	public String getName() {
		return customer.getName();
	}

	public Date getDate() {
		return transactionDate;
	}

	public List<Product> getProducts() {
		return products;
	}

//65070501074 Napat Sinjindawong Group 5


//	public String productBreakDown() {
//		return "list of products for reporting";
//	}
//
//	public void prepareInvoice() {
//		System.out.println("invoice prepared...");
//	}
//
//	public void chargeCustomer() {
//		System.out.println("charged the customer");
//	}
}
