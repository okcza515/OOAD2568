
public class ReportGenerator {
	private CustomerTransaction transaction;

	public ReportGenerator(CustomerTransaction transaction){
		this.transaction = transaction;
	}
	
	public void generateReport(){
		System.out.println("Transaction Report:\n  Customer: "+ transaction.getCustomerName() + 
			"\n  Items on: " + transaction.getTransactionDate());
	
		System.out.println("\n  Purchase " + transaction.getProducts().size() + " Products" + 
			"\n  List of Product:");
		for(Product product : transaction.getProducts()){
			System.out.println("  Product ID: " + product.getProductId() + " Product Name: " + product.getProductName());
		}	
	}

	


}

//65070501053