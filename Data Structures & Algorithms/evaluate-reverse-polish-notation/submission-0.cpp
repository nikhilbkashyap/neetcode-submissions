class Solution {
public:
    int evalRPN(vector<string>& tokens) {
        while (tokens.size() > 1) {
            for (int i = 0; i < tokens.size(); i++) 
            {
                if (tokens[i] == "+" || tokens[i] == "-" || tokens[i] == "*" || tokens[i] == "/") 
                {
                    int a= stoi(tokens[i - 2]);
                    int b= stoi(tokens[i - 1]);
                    int result = 0;
                    if (tokens[i] == "+")
                        result = a + b;
                    else if (tokens[i] == "-")
                        result = a - b;
                    else if (tokens[i] == "*")
                        result = a * b;
                    else if (tokens[i] == "/")
                        result = a / b;

                    vector<string> newTokens;
                    
                    for (int j = 0; j < i - 2; ++j)
                        newTokens.push_back(tokens[j]);
                    
                    newTokens.push_back(to_string(result));
                    
                    for (int j = i + 1; j < tokens.size(); ++j)
                        newTokens.push_back(tokens[j]);
                    
                    tokens = newTokens;
                    break;
                }
            }
        }
        return stoi(tokens[0]);
    }
};
