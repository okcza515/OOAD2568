// 1. Customer related interfaces and classes\
//Chayaphon Chaisangkha 65070503409

import java.math.BigDecimal;
import java.time.LocalDateTime;
import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.atomic.AtomicLong;

interface Customer {
    String getId();
    String getName();
}

class CustomerImpl implements Customer {
    private String id;
    private String name;

    public CustomerImpl(String id, String name) {
        this.id = id;
        this.name = name;
    }

    @Override
    public String getId() {
        return id;
    }

    @Override
    public String getName() {
        return name;
    }
}

// 2. Product related interfaces and classes
//Bew
interface Product {
    String getId();
    String getName();
    String getBreakdown();
    BigDecimal getPrice();
}

class ProductImpl implements Product {
    private String id;
    private String name;
    private String breakdown;
    private BigDecimal price;

    public ProductImpl(String id, String name, String breakdown, BigDecimal price) {
        this.id = id;
        this.name = name;
        this.breakdown = breakdown;
        this.price = price;
    }

    @Override
    public String getId() { return id; }

    @Override
    public String getName() { return name; }

    @Override
    public String getBreakdown() { return breakdown; }

    @Override
    public BigDecimal getPrice() { return price; }
}

// 4. Transaction Generator interface and implementation
// Sawitt Ngamvilaisiriwong 65070503469
interface TransactionGenerator {
    Transaction generateTransaction(Customer customer, Product product) throws TransactionException;
}

class TransactionGeneratorImpl implements TransactionGenerator {
    private static final String TRANSACTION_ID_PREFIX = "TRX";
    private final AtomicLong transactionCounter = new AtomicLong(0);

    @Override
    public Transaction generateTransaction(Customer customer, Product product) throws TransactionException {
        String transactionId = TRANSACTION_ID_PREFIX + transactionCounter.incrementAndGet();
        return new CustomerTransaction(transactionId, customer, product);
    }
}

// 5. Accounts Receivable interface and implementation
interface AccountsReceivable {
    void recordTransaction(Transaction transaction) throws AccountingException;
    BigDecimal getCustomerBalance(Customer customer);
    List<Transaction> getCustomerTransactions(Customer customer);
}

class AccountsReceivableImpl implements AccountsReceivable {
    private Map<String, List<Transaction>> customerTransactions = new ConcurrentHashMap<>();
    private Map<String, BigDecimal> customerBalances = new ConcurrentHashMap<>();

    @Override
    public void recordTransaction(Transaction transaction) throws AccountingException {
        String customerId = transaction.getCustomer().getId();
        customerTransactions.computeIfAbsent(customerId, k -> new ArrayList<>())
                          .add(transaction);
        
        customerBalances.compute(customerId, (k, v) -> 
            (v == null ? BigDecimal.ZERO : v).add(transaction.getAmount()));
    }

    @Override
    public BigDecimal getCustomerBalance(Customer customer) {
        return customerBalances.getOrDefault(customer.getId(), BigDecimal.ZERO);
    }

    @Override
    public List<Transaction> getCustomerTransactions(Customer customer) {
        return customerTransactions.getOrDefault(customer.getId(),Collections.emptyList());
    }
}

// 6. Service layer to orchestrate the operations
//Chanwat Limpanatewin 65070503445
class TransactionService {
    private final TransactionGenerator transactionGenerator;
    private final AccountsReceivable accountsReceivable;

    public TransactionService(TransactionGenerator transactionGenerator, 
                            AccountsReceivable accountsReceivable) {
        this.transactionGenerator = transactionGenerator;
        this.accountsReceivable = accountsReceivable;
    }

    public void processTransaction(Customer customer, Product product) 
            throws TransactionException, AccountingException {
        Transaction transaction = transactionGenerator.generateTransaction(customer, product);
        transaction.execute();
        accountsReceivable.recordTransaction(transaction);
    }

    public BigDecimal getCustomerBalance(Customer customer) {
        return accountsReceivable.getCustomerBalance(customer);
    }
}

// 7. Custom exceptions
class TransactionException extends Exception {
    public TransactionException(String message) {
        super(message);
    }
}

class AccountingException extends Exception {
    public AccountingException(String message) {
        super(message);
    }
}

// 3. Transaction related interfaces and classes
// 65070503466 Warapol Pratumta
interface Transaction {
    String getTransactionId();
    Customer getCustomer();
    Product getProduct();
    BigDecimal getAmount();
    LocalDateTime getTransactionDate();
    void execute() throws TransactionException;
}

class CustomerTransaction implements Transaction {
    private String transactionId;
    private Customer customer;
    private Product product;
    private BigDecimal amount;
    private LocalDateTime transactionDate;

    public CustomerTransaction(String transactionId, Customer customer, Product product) {
        this.transactionId = transactionId;
        this.customer = customer;
        this.product = product;
        this.amount = product.getPrice();
        this.transactionDate = LocalDateTime.now();
    }

    @Override
    public String getTransactionId() { return transactionId; }

    @Override
    public Customer getCustomer() { return customer; }

    @Override
    public Product getProduct() { return product; }

    @Override
    public BigDecimal getAmount() { return amount; }

    @Override
    public LocalDateTime getTransactionDate() { return transactionDate; }

    @Override
    public void execute() throws TransactionException {
        // Implementation of transaction execution
    }
}