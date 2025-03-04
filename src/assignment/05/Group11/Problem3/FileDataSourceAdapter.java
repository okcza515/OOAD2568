public class FileDataSourceAdapter implements DataSource {
    private FileDataSource fileDataSource;
    
    public FileDataSourceAdapter(String name) {
        this.fileDataSource = new FileDataSource(name);
    }
    
    @Override
    public void writeData(String data) {
        fileDataSource.writeData(data);
    }
    
    @Override
    public String readData() {
        return fileDataSource.readData();
    }
} 
// 65070503412 Chitsanupong Jateassavapirom