// 65070503403 Kritsanaphong Thaworana
import java.util.Date;
import java.util.List;

public class CustomerTransaction {
    private Customer customer;
    private List<Product> products;
    private Date date;

    public CustomerTransaction(Customer customer, List<Product> products) {
        if (customer == null || products == null) {
            throw new IllegalArgumentException("Customer and products cannot be null.");
        }
        this.customer = customer;
        this.products = products;
        this.date = new Date(); // Set current date
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
            breakdown.append(" - ").append(product.getProductName())
                     .append(" (ID: ").append(product.getProductId()).append(")\n");
        }
        return breakdown.toString();
    }
}
