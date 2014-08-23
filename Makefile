BINARIES = saygaijin
BLDDIR = build
SRCDIR = .

all : $(BINARIES)

$(BINARIES) : %: $(BLDDIR)/%

$(BLDDIR)/saygaijin:
	mkdir -p $(BLDDIR)
	cd $(SRCDIR) && go build -o $(abspath $@)

clean:
	rm -fr $(BLDDIR)
