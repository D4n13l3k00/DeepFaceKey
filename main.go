package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/anacrolix/torrent"
)

func main() {

	a := app.New()
	w := a.NewWindow("DeepFaceKey")
	c, _ := torrent.NewClient(nil)
	t, _ := c.AddMagnet("magnet:?xt=urn:btih:e7ffdcb4ada863de9504f2a741f924dcd56ab84a&amp;dn=DeepFaceLab&amp;tr=udp%3a%2f%2ftracker.dler.com%3a6969%2fannounce&amp;tr=http%3a%2f%2ftracker3.itzmx.com%3a6961%2fannounce&amp;tr=udp%3a%2f%2ftracker.blacksparrowmedia.net%3a6969%2fannounce&amp;tr=http%3a%2f%2ftracker2.itzmx.com%3a6961%2fannounce&amp;tr=udp%3a%2f%2fengplus.ru%3a6969%2fannounce&amp;tr=udp%3a%2f%2ftracker.torrent.eu.org%3a451%2fannounce&amp;tr=http%3a%2f%2ft.nyaatracker.com%3a80%2fannounce&amp;tr=udp%3a%2f%2ftracker.zerobytes.xyz%3a1337%2fannounce&amp;tr=udp%3a%2f%2ftracker0.ufibox.com%3a6969%2fannounce&amp;tr=udp%3a%2f%2ftracker2.dler.org%3a80%2fannounce&amp;tr=http%3a%2f%2fretracker.sevstar.net%3a2710%2fannounce&amp;tr=udp%3a%2f%2finferno.demonoid.is%3a3391%2fannounce&amp;tr=https%3a%2f%2ftracker.lilithraws.cf%3a443%2fannounce&amp;tr=udp%3a%2f%2fbt2.archive.org%3a6969%2fannounce&amp;tr=udp%3a%2f%2fopen.stealth.si%3a80%2fannounce&amp;tr=udp%3a%2f%2fcode2chicken.nl%3a6969%2fannounce&amp;tr=udp%3a%2f%2fmail.realliferpg.de%3a6969%2fannounce&amp;tr=http%3a%2f%2ftracker.files.fm%3a6969%2fannounce&amp;tr=udp%3a%2f%2ftracker.0x.tf%3a6969%2fannounce&amp;tr=udp%3a%2f%2fdiscord.heihachi.pw%3a6969%2fannounce&amp;tr=udp%3a%2f%2fmts.tvbit.co%3a6969%2fannounce&amp;tr=http%3a%2f%2fvps02.net.orel.ru%3a80%2fannounce&amp;tr=udp%3a%2f%2fwww.torrent.eu.org%3a451%2fannounce&amp;tr=udp%3a%2f%2fbt1.archive.org%3a6969%2fannounce&amp;tr=udp%3a%2f%2fadmin.videoenpoche.info%3a6969%2fannounce&amp;tr=udp%3a%2f%2ftracker.altrosky.nl%3a6969%2fannounce&amp;tr=udp%3a%2f%2fmovies.zsw.ca%3a6969%2fannounce&amp;tr=https%3a%2f%2ftracker.tamersunion.org%3a443%2fannounce&amp;tr=udp%3a%2f%2fretracker.lanta-net.ru%3a2710%2fannounce&amp;tr=https%3a%2f%2ftracker.nitrix.me%3a443%2fannounce")
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(680, 400))

	var dl_files []*torrent.File

	log_box := NewLogger()

	progress_bar := widget.NewProgressBar()
	progress_bar.Hide()

	select_file := widget.NewSelect([]string{"Please, wait..."}, func(_ string) {})
	select_file.Hide()
	select_file.OnChanged = func(s string) {
		select_file.Hide()
		log_box.log(fmt.Sprintf("Selected `%s`", s))
		log_box.log("Starting download...")
		go func() {
			file := dl_files[select_file.SelectedIndex()]
			log_box.log(fmt.Sprintf("Downloading `%s`", file.DisplayPath()))
			file.Download()
			progress_bar.TextFormatter = func() string {
				val := math.Floor(float64(file.BytesCompleted()) / float64(file.Length()) * 100)
				progress_bar.Value = val / 100
				return fmt.Sprintf("%d%% | Peers: %d | Seeders: %d", int64(val), t.Stats().TotalPeers, t.Stats().ConnectedSeeders)

			}
			progress_bar.Show()
			for {
				progress_bar.Refresh()
				if file.BytesCompleted() == file.Length() {
					progress_bar.Hide()
					log_box.log("Download complete!")
					progress_bar.Hide()
					DeleteTempFiles()
					break
				}
			}
		}()
	}

	w.SetContent(container.NewVBox(
		progress_bar,
		select_file,
		log_box,
	))

	go func() {
		log_box.log("Using magnet link...")
		log_box.log("Fetching info...")
		<-t.GotInfo()
		log_box.log("Fetching files...")
		dl_files = t.Files()
		chose_files_names := []string{}
		for _, file := range dl_files {
			chose_files_names = append(chose_files_names, file.DisplayPath())
		}
		log_box.log("Select needed file...")
		select_file.Options = chose_files_names
		select_file.Show()
	}()

	w.ShowAndRun()
}

func DeleteTempFiles() {
	files, err := filepath.Glob(".torrent.db*")
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		if err := os.Remove(f); err != nil {
			panic(err)
		}
	}
}
