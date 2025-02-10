//65070501008
public class ReportGenerator {
	
	private CustomerTransaction transactionObject;

	public ReportGenerator(CustomerTransaction transactionObject){
		this.transactionObject = transactionObject;
	}
	
	public void generateReport(){
		System.out.println("===== Transaction Report =====");
        System.out.println("Customer: " + transactionObject.getName());
        System.out.println("Date: " + transactionObject.getDate());
        System.out.println(transactionObject.productBreakDown());
	}

}
