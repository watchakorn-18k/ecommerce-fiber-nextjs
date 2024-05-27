interface PostProductData {
    product_name: string;
    quantity: number;
    price: number;
    product_id: string;
}

export default async function addProduct({ product_name, quantity, price, product_id }: PostProductData): Promise<any> {
    const data: PostProductData = {
        product_name,
        quantity,
        price,
        product_id
    };

    try {
        const response = await fetch(`${process.env.BACKEND_URL}/api/v1/product/add_product`, {
            method: "POST",
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(data)
        });

        const responseText = await response.text();

        if (!response.ok) {
            console.error('Server response (error):', responseText);
            let errorDetails;
            try {
                errorDetails = JSON.parse(responseText);
            } catch (e) {
                throw new Error(`Failed to fetch product: ${response.statusText}`);
            }
            throw new Error(`Failed to fetch product: ${errorDetails.message || response.statusText}`);
        }

        console.log('Server response (success):', responseText);
        let responseData;
        try {
            responseData = JSON.parse(responseText);
        } catch (e) {
            throw new Error('Invalid JSON response from server');
        }

        return responseData;
    } catch (error) {
        console.error('Error:', error);
        throw error;
    }
}