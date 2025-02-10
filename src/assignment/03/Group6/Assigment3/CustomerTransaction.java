//65070501008
import java.util.Date;
import java.util.List;

public class CustomerTransaction {

	private List<Product> products;
	private Customer customer;
	private Date date;

	public CustomerTransaction(Customer customer, List<Product> products) {
		this.products = products;
		this.customer = customer;
		this.date = new Date();
	}

	public String getName() {
        return customer.getName();
    }

	public Date getDate() {
		return date;
	}

	public List<Product> getProducts() {
		return products;
	}

	public String productBreakDown() {
		if (products.isEmpty()) {
			return "No products in transaction.";
		}
		
		StringBuilder breakdown = new StringBuilder("Products in Transaction:\n");
		for (Product product : products) {
			breakdown.append(String.format(" - ID: %d | Name: %s\n", 
					product.getProductId(), product.getProductName()));
		}
		
		return breakdown.toString();
	}
	
	
}
