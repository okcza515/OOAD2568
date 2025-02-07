import java.util.Date;
import java.util.List;

interface InvoiceHandler {
	void prepareInvoice();
}

class InvoiceService implements InvoiceHandler {
	@Override
	public void prepareInvoice() {
		System.out.println("Invoice has been prepared.");
	}
}

public class CustomerTransaction {

	private List<Product> products;
	private Customer customer;
	private Date transactionDate;

	public CustomerTransaction(Customer customer, List<Product> products) {
		this.products = products;
		this.customer = customer;
		this.transactionDate = new Date();
	}

	public String getCustomerName() {
		return customer.getName();
	}

	public Date getTransactionDate() {
		return transactionDate;
	}

	public List<Product> getProducts() {
		return products;
	}

	public void prepareInvoice() {
		System.out.println("Invoice prepared for " + getCustomerName());
	}
}

//65070501053