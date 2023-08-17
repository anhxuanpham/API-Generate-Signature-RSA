# SignWithRSA

1. Chuẩn bị file chứa private key
2. Đổi tên file thành private_key.pem
3. Đưa dữ liệu cần ký vào biến datatoSign

#REQ

{
    "data": "merchant_code=BEANBAKERY&account_name=NGUYEN VAN A&map_id=052097000111&map_type=CCCD&account_type=M&bank_code=TCB"
}


#RES

{
    "signature": "TRziwp1cbL4Gd0b4bInrmLUSsBh7odM2sM3opAn9YSvByK/eV++ofkRcu9YZdEMBG2aPFnhsV1R4JPMwtAijFdyLp376E0iSQW7ArsCJOPalOU0yDlS5I6wXSb70FosePMqnuUoA0+lPJqz981CB2rnomyx0V5+JEiCfW5laxYq1w2G9GLu3BealmqkoxTQdwUyQXkDHSjtBntPc81T3P0aQdIxYbJ2gQVjJA4oJGV06RlMxGUf8ODGqWMikE9J5FBNL7jUlA4C0Qljh+Za14v4wDs8mR0/7bk8ZivZS8OJGPKCE2CmQe/uvp6BvjeSp1KgQ17vlA6l53u5gLjICOQ=="
}


# Docker
docker-compose up --build -d
