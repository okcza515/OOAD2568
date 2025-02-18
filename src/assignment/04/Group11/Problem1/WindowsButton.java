import javax.swing.*;
import java.awt.*;
import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;

public class WindowsButton implements Button {  // Add "implements Button" here
    JPanel panel = new JPanel();
    JFrame frame = new JFrame();
    JButton button;

    public WindowsButton() {
        button = new JButton("Exit");
    }

    // @Override  // Add @Override annotations
    public void render() {
        frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        JLabel label = new JLabel("Hello Wit03!");
        label.setOpaque(true);
        label.setBackground(new Color(235, 233, 126));
        label.setFont(new Font("Dialog", Font.BOLD, 44));
        label.setHorizontalAlignment(SwingConstants.CENTER);
        panel.setLayout(new FlowLayout(FlowLayout.CENTER));
        frame.getContentPane().add(panel);
        panel.add(label);
        panel.add(button);
        
        frame.setSize(320, 200);
        frame.setVisible(true);
        onClick();
    }
    
    // @Override  // Add @Override annotation
    public void onClick() {
        button.addActionListener(new ActionListener() {
            public void actionPerformed(ActionEvent e) {
                frame.setVisible(false);
                System.exit(0);
            }
        });
    }
}