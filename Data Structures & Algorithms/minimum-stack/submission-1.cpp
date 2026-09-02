class MinStack {
public:
        stack<int> s;
    MinStack() {
        
    }
    
    void push(int val) {
        s.push(val);
    }
    
    void pop() {
        s.pop();
    }
    
    int top() {
        return s.top();
    }
    
    int getMin() {
        stack<int> a;
        int miN = s.top();
        while(s.size()){
            miN = min(miN, s.top());
            a.push(s.top());
            s.pop();
        }
        while(a.size()){
            s.push(a.top());
            a.pop();
        }
        return miN;
    }
};
