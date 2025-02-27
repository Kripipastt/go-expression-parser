import axios from 'axios'
import {Expression} from "./interface";
// import {API_URL} from "../config"

class ExpressionService {
    private static API = "http://localhost:8080"//API_URL;

    public static async sendExpression(expression: string): Promise<void> {
        const response = await axios.post(this.API + "/api/v1/calculate", {expression}).catch(err => {
            if (err.response) {
                return {data: []}
            }
        });
        return response.data
    }

    public static async getAllExpressions(): Promise<Expression[]> {
        const response = await axios.get(this.API + "/api/v1/expressions")
        return response.data.expressions
    }

    public static async getExpression(expression_id: string): Promise<Expression> {
        const response = await axios.get(this.API + "/api/v1/expressions/" + expression_id);
        return response.data;
    }
}

export default ExpressionService;