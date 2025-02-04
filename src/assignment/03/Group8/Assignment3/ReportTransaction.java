import java.util.Date;
import java.util.List;

public class ReportTransaction {

	private List<Product> products;
	private Customer customer;

	public ReportTransaction(Customer customer, List<Product> products) {
		this.products = products;
		this.customer = customer;
	}

	public String getName() {
		return customer.getName();
	}

	public Date getDate() {
		return new Date();
	}

	public String productBreakDown() {
		return "list of products for reporting";
	}

	
}
