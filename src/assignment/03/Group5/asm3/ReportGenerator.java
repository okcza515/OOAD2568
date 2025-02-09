public class ReportGenerator {

	public void generateReport(CustomerTransaction transaction) {
		System.out.println("Transaction Report:");
		System.out.println("Customer: " + transaction.getName());
		System.out.println("Date: " + transaction.getDate());
		System.out.println("Products Purchased:");
		for (Product product : transaction.getProducts()) {
			System.out.println("- " + product.getProductName());
		}
	}
}

//65070501074 Napat Sinjindawong Group 5