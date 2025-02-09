
public class ReportGenerator {
	
	private CustomerTransaction transactionObject;
	
	public void generateReport(){
		System.out.println(transactionObject.getName()+" "+transactionObject.productBreakDown()+" "+transactionObject.getDate());
	}

}
