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
        StringBuilder breakdown = new StringBuilder("Product Breakdown:\n");
        for (Product product : products) {
            breakdown.append("ID: ").append(product.getProductId())
                     .append(", Name: ").append(product.getProductName())
                     .append("\n");
        }
        return breakdown.toString();
	}
}
