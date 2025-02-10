// 65070501019 Natlada Simasathien
public class ReportGenerator {
	
	private CustomerTransaction transactionObject;

	public ReportGenerator(CustomerTransaction transactionObject) {
		this.transactionObject = transactionObject;
	}
	
	public void generateReport(){
		System.out.println(transactionObject.getName()+"\n"+transactionObject.productBreakDown()+"\n"+transactionObject.getDate());
	}
	
}