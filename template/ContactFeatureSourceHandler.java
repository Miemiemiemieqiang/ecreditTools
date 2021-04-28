package com.ecreditpal.vodka.variable.handler.source;

import com.ecreditpal.vodka.model.entity.backdata.KsModelPhoneContactVarInfo;
import com.ecreditpal.vodka.variable.handler.Invoke;
import com.ecreditpal.vodka.variable.handler.Source;
import com.ecreditpal.vodka.variable.handler.SourceHandler;
import org.springframework.util.CollectionUtils;

import java.util.List;
import java.util.Objects;


/**
 * Created by {{.Author}} on {{.Datetime}}
 */
@Source(value = "{{.Lower}}", alias = "{{.Lower}}")
public class {{.Upper}}SourceHandler extends SourceHandler {

    @Invoke
    public Integer getSource(List<KsModelPhoneContactVarInfo> ksModelPhoneContactVarList) {
        if (CollectionUtils.isEmpty(ksModelPhoneContactVarList)
                || ksModelPhoneContactVarList.stream().map(KsModelPhoneContactVarInfo::{{.Extra}}).allMatch(Objects::isNull)) {
            return null;
        }
        return ksModelPhoneContactVarList.stream().filter(Objects::nonNull).mapToInt(KsModelPhoneContactVarInfo::{{.Extra}}).sum();
    }

}
