package main

import "github.com/rs/zerolog/log"

func main() {

	err := Gen_AFC_DS01_Pain013_001_07("docmapper/defs/AFC-DS01-pain.013.001.07.yml")
	if err != nil {
		log.Error().Err(err).Msg("Creating Pain013")
	}

	err = Gen_AFC_DS10_Camt055_001_08("docmapper/defs/AFC-DS10-camt.055.001.08.yml")
	if err != nil {
		log.Error().Err(err).Msg("Creating Camt055")
	}

	err = Gen_AFC_DS15_Pacs028_001_03("docmapper/defs/AFC-DS15-pacs.028.001.03.yml")
	if err != nil {
		log.Error().Err(err).Msg("Creating Pacs028")
	}

	log.Info().Msg("done")
}
