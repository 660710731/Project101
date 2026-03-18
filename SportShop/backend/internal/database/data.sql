insert into categories (name) values ('Football'), ('Basketball'), ('Tennis'), ('Running'), ('Swimming');
insert into products (name, description, price, stock, category_id, image_url) values 
('Football', 'High-quality football for all weather conditions', 29.99, 100, 1, 'https://avecsport.com/images/blogs/hyperbluefootball.jpg'),
('Basketball', 'Durable basketball suitable for indoor and outdoor play', 24.99, 150, 2, 'https://www.shutterstock.com/image-photo/closeup-basketball-on-indoor-court-600nw-2665867459.jpg'),
('Tennis Racket', 'Lightweight tennis racket with excellent grip', 89.99, 50, 3, 'https://images.prodirectsport.com/productimages/Main/1031501_Main_2024768.jpg'),
('Running Shoes', 'Comfortable running shoes with great support', 119.99, 200, 4, 'https://d1i8d6ce8wwubf.cloudfront.net/photos/28/73/408849_7358_XL.jpg'),
('Swimming Goggles', 'Anti-fog swimming goggles with UV protection', 19.99, 300, 5, 'https://t4.ftcdn.net/jpg/03/02/14/25/360_F_302142545_Z0qAD5aqjUjaEde8wmShwKlZMhMNoeiE.jpg');
insert into users (username, password, role) values 
('employee1', 'password123', 'employee'),
('customer1', 'password456', 'customer') ON CONFLICT (username) DO NOTHING;