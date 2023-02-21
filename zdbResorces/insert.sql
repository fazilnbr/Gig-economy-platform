\COPY users (user_name,password,user_type,verification,status) from 'logins.txt' WITH  DELIMITER ',';

\COPY profiles (login_id,name,gender,date_of_birth,house_name,place,post,pin,contact_number,email_id,photo) from 'profiles.txt' WITH  DELIMITER ',';

\COPY categories (category) from 'categories.txt' WITH  DELIMITER ',';

\COPY jobs (category_id,id_worker,wage,description) from 'jobs.txt' WITH  DELIMITER ',';

\COPY addresses (user_id,house_name,place,city,post,pin,phone) from 'address.txt' WITH  DELIMITER ',';

