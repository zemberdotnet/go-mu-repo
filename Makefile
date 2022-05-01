main:
	go build -o gum


clean:
	rm .gum
	rm -rf blog
	rm -rf graph-cisco-secure-endpoint
	rm -rf graph-rumble

run: main clean
	./gum clone https://github.com/jupiterone/graph-rumble
	./gum register graph-rumble
	./gum clone https://github.com/jupiterone/graph-cisco-secure-endpoint
	./gum register graph-cisco-secure-endpoint
	./gum pull
	./gum checkout -b test
	./gum sh yarn
	./gum sh touch test.txt
	./gum add -A
	./gum commit -m "test"