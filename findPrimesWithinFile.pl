$fileToBeRead = "result.txt";
$fileToBeWrite = "primeResult.txt";
open(readFile,'<',$fileToBeRead) or die "Sorry, File to be read is not found.";
open(writeFile,'>',$fileToBeWrite) or die "Sorry, File to be write is not found.";
$lastLineCursor=0;
while(<readFile>)
{
    @fragmentedLine = split(' ', $_);
    if($fragmentedLine[4] eq "prime"){
        if($lastLineCursor != 0){ # don't put cursor to last line
            print writeFile "\n";
        }
        print writeFile $fragmentedLine[0]; 
        $lastLineCursor++;
    }
    
}
close;
close;