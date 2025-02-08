
public class ReportGenerator {
	
	private CustomerTransaction transactionObject;

	public ReportGenerator(CustomerTransaction transactionObject){
		this.transactionObject = transactionObject;
	}
	
	public void generateReport(){
		System.out.println("Report:");
		System.out.println(transactionObject.getName()+"\n"+transactionObject.productBreakDown()+transactionObject.getDate());
	}

}
