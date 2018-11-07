import java.util.Properties;

public class PrintSystemInformation {
    public void printInformation() {
        Properties p = System.getProperties();
        p.list(System.out);

        System.out.println("-----Memory Usage:");

        Runtime rt = Runtime.getRuntime();
        System.out.println("Total Memory = " + rt.totalMemory() );
    }
}
