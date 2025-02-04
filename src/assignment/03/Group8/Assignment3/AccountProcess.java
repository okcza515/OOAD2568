import java.util.List;

public class AccountProcess {
    
    private List<Product> products;
	private Customer customer;

	public AccountProcess(Customer customer, List<Product> products) {
		this.products = products;
		this.customer = customer;
	}

    public void prepareInvoice() {
		System.out.println("invoice prepared...");
	}

	public void chargeCustomer() {
		System.out.println("charged the customer");
	}
}
