/*
Real-time Online/Offline Charging System (OCS) for Telecom & ISP environments
Copyright (C) ITsysCOM GmbH

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>
*/

package config

import (
	"github.com/cgrates/cgrates/utils"
)

type RateSCfg struct {
	Enabled                 bool
	IndexedSelects          bool
	StringIndexedFields     *[]string
	PrefixIndexedFields     *[]string
	SuffixIndexedFields     *[]string
	NestedFields            bool
	RateIndexedSelects      bool
	RateStringIndexedFields *[]string
	RatePrefixIndexedFields *[]string
	RateSuffixIndexedFields *[]string
	RateNestedFields        bool
}

func (rCfg *RateSCfg) loadFromJsonCfg(jsnCfg *RateSJsonCfg) (err error) {
	if jsnCfg == nil {
		return
	}
	if jsnCfg.Enabled != nil {
		rCfg.Enabled = *jsnCfg.Enabled
	}
	if jsnCfg.Indexed_selects != nil {
		rCfg.IndexedSelects = *jsnCfg.Indexed_selects
	}
	if jsnCfg.String_indexed_fields != nil {
		sif := make([]string, len(*jsnCfg.String_indexed_fields))
		for i, fID := range *jsnCfg.String_indexed_fields {
			sif[i] = fID
		}
		rCfg.StringIndexedFields = &sif
	}
	if jsnCfg.Prefix_indexed_fields != nil {
		pif := make([]string, len(*jsnCfg.Prefix_indexed_fields))
		for i, fID := range *jsnCfg.Prefix_indexed_fields {
			pif[i] = fID
		}
		rCfg.PrefixIndexedFields = &pif
	}
	if jsnCfg.Suffix_indexed_fields != nil {
		sif := make([]string, len(*jsnCfg.Suffix_indexed_fields))
		for i, fID := range *jsnCfg.Suffix_indexed_fields {
			sif[i] = fID
		}
		rCfg.SuffixIndexedFields = &sif
	}
	if jsnCfg.Nested_fields != nil {
		rCfg.NestedFields = *jsnCfg.Nested_fields
	}

	if jsnCfg.Rate_indexed_selects != nil {
		rCfg.RateIndexedSelects = *jsnCfg.Rate_indexed_selects
	}
	if jsnCfg.Rate_string_indexed_fields != nil {
		sif := make([]string, len(*jsnCfg.Rate_string_indexed_fields))
		for i, fID := range *jsnCfg.Rate_string_indexed_fields {
			sif[i] = fID
		}
		rCfg.RateStringIndexedFields = &sif
	}
	if jsnCfg.Rate_prefix_indexed_fields != nil {
		pif := make([]string, len(*jsnCfg.Rate_prefix_indexed_fields))
		for i, fID := range *jsnCfg.Rate_prefix_indexed_fields {
			pif[i] = fID
		}
		rCfg.RatePrefixIndexedFields = &pif
	}
	if jsnCfg.Rate_suffix_indexed_fields != nil {
		sif := make([]string, len(*jsnCfg.Rate_suffix_indexed_fields))
		for i, fID := range *jsnCfg.Rate_suffix_indexed_fields {
			sif[i] = fID
		}
		rCfg.RateSuffixIndexedFields = &sif
	}
	if jsnCfg.Rate_nested_fields != nil {
		rCfg.RateNestedFields = *jsnCfg.Rate_nested_fields
	}
	return
}

func (rCfg *RateSCfg) AsMapInterface() map[string]interface{} {
	stringIndexedFields := []string{}
	if rCfg.StringIndexedFields != nil {
		stringIndexedFields = make([]string, len(*rCfg.StringIndexedFields))
		for i, item := range *rCfg.StringIndexedFields {
			stringIndexedFields[i] = item
		}
	}
	prefixIndexedFields := []string{}
	if rCfg.PrefixIndexedFields != nil {
		prefixIndexedFields = make([]string, len(*rCfg.PrefixIndexedFields))
		for i, item := range *rCfg.PrefixIndexedFields {
			prefixIndexedFields[i] = item
		}
	}
	rateStringIndexedFields := []string{}
	if rCfg.RateStringIndexedFields != nil {
		rateStringIndexedFields = make([]string, len(*rCfg.RateStringIndexedFields))
		for i, item := range *rCfg.RateStringIndexedFields {
			rateStringIndexedFields[i] = item
		}
	}
	ratePrefixIndexedFields := []string{}
	if rCfg.RatePrefixIndexedFields != nil {
		ratePrefixIndexedFields = make([]string, len(*rCfg.RatePrefixIndexedFields))
		for i, item := range *rCfg.RatePrefixIndexedFields {
			ratePrefixIndexedFields[i] = item
		}
	}
	return map[string]interface{}{
		utils.EnabledCfg:                 rCfg.Enabled,
		utils.IndexedSelectsCfg:          rCfg.IndexedSelects,
		utils.StringIndexedFieldsCfg:     stringIndexedFields,
		utils.PrefixIndexedFieldsCfg:     prefixIndexedFields,
		utils.NestedFieldsCfg:            rCfg.NestedFields,
		utils.RateIndexedSelectsCfg:      rCfg.RateIndexedSelects,
		utils.RateStringIndexedFieldsCfg: rateStringIndexedFields,
		utils.RatePrefixIndexedFieldsCfg: ratePrefixIndexedFields,
		utils.RateNestedFieldsCfg:        rCfg.RateNestedFields,
	}
}