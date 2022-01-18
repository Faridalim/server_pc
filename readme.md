Proses di client :

1. Web app ujian fetch get ke localhost untuk mengupdate kondisi PC sedang digunakan atau tidak per 5 detik
2. App client nge post ke server pc setiap 5 detik mengirim kondisi PC sedang in use atau tidak
3. Jika request terakhir berlalu s.d lebih dari 15 detik maka app client akan mengirim kondisi PC in use = false ke server pc
4. Display akan mengupdate setiap 5 detik. Jika response last update lebih dari 15 detik yang lalu maka pc dianggap off